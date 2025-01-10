#include <pbc/pbc.h>
#include <iostream>
#include <vector>
#include <cstring>
#include <sstream>
#include <iomanip>
#include <algorithm>

// 定义BLS签名结构
struct BLSSignature {
    element_t sig; // 签名
};

// 初始化BLS签名系统
void init_bls(pairing_t pairing, element_t g, element_t sk, element_t pk) {
    // 选择随机私钥
    element_init_Zr(sk, pairing);
    element_random(sk);

    // 计算公钥
    element_init_G1(pk, pairing);
    element_mul_zn(pk, g, sk);
}

// 生成BLS签名
void sign_bls(pairing_t pairing, element_t g, element_t sk, const char* message, BLSSignature& signature) {
    element_t h;
    element_init_G1(h, pairing);
    element_from_hash(h, (void*)message, strlen(message));

    element_init_G1(signature.sig, pairing);
    element_mul_zn(signature.sig, h, sk);
}

// 将 element_t 序列化为字符串
std::string element_to_string(element_t elem) {
    int len = element_length_in_bytes(elem);
    std::vector<unsigned char> data(len);
    element_to_bytes(data.data(), elem);

    std::ostringstream oss;
    oss << std::hex;
    for (int i = 0; i < len; ++i) {
        oss << std::setw(2) << std::setfill('0') << (int)data[i];
    }
    return oss.str();
}

// 将字符串反序列化为 element_t
void string_to_element(element_t& elem, const std::string& str) {
    std::vector<unsigned char> data(str.length() / 2);
    for (size_t i = 0; i < data.size(); ++i) {
        std::string byte = str.substr(2 * i, 2);
        data[i] = static_cast<unsigned char>(std::stoul(byte, nullptr, 16));
    }
    element_from_bytes(elem, data.data());
}

// 将 BLSSignature 序列化为字符串
std::string signature_to_string(BLSSignature& sig) {
    return element_to_string(sig.sig);
}

// 将字符串反序列化为 BLSSignature
void string_to_signature(BLSSignature& sig, const std::string& str, pairing_t pairing) {
    element_init_G1(sig.sig, pairing);
    string_to_element(sig.sig, str);
}

// 聚合BLS签名
void aggregate_signatures(pairing_t pairing, std::vector<BLSSignature>& signatures, BLSSignature& aggregated_signature) {
    element_init_G1(aggregated_signature.sig, pairing);
    element_set0(aggregated_signature.sig);

    for (auto& sig : signatures) {
        element_add(aggregated_signature.sig, aggregated_signature.sig, sig.sig);
    }
}

// 分割字符串
std::vector<std::string> split_string(const std::string& str, char delimiter) {
    std::vector<std::string> tokens;
    std::string token;
    std::istringstream token_stream(str);
    while (std::getline(token_stream, token, delimiter)) {
        tokens.push_back(token);
    }
    return tokens;
}

// 验证聚合签名（接受拼接的公钥字符串）
bool verify_aggregated_signature(
        pairing_t pairing,
        element_t g,
        const std::string& pk_string_combined,
        const std::vector<const char*>& messages,
        const std::string& aggregated_signature_string
) {
    // 分割公钥字符串
    std::vector<std::string> pk_strings = split_string(pk_string_combined, ',');

    // 反序列化公钥
    std::vector<element_t> pks(pk_strings.size());
    for (size_t i = 0; i < pk_strings.size(); ++i) {
        element_init_G1(pks[i], pairing);
        string_to_element(pks[i], pk_strings[i]);
    }

    // 反序列化聚合签名
    BLSSignature aggregated_signature;
    string_to_signature(aggregated_signature, aggregated_signature_string, pairing);

    // 验证聚合签名
    element_t left, right, h;
    element_init_GT(left, pairing);
    element_init_GT(right, pairing);
    element_init_G1(h, pairing);

    pairing_apply(left, aggregated_signature.sig, g, pairing);

    element_t prod;
    element_init_GT(prod, pairing);
    element_set1(prod);

    for (size_t i = 0; i < pks.size(); ++i) {
        element_from_hash(h, (void*)messages[i], strlen(messages[i]));
        pairing_apply(right, h, pks[i], pairing);
        element_mul(prod, prod, right);
    }

    // 清理反序列化的元素
    for (auto& pk : pks) {
        element_clear(pk);
    }
    element_clear(aggregated_signature.sig);

    return element_cmp(left, prod) == 0;
}

int main() {
    // 硬编码的 a.param 参数
    const char* fixed_param_str =
            "type a\n"
            "q 8780710799663312522437781984754049815806883199414208211028653399266475630880222957078625179422662221423155858769582317459277713367317481324925129998224791\n"
            "h 12016012264891146079388821366740534204802954401251311822919615131047207289359704531102844802183906537786776\n"
            "r 730750818665451621361119245571504901405976559617\n"
            "exp2 159\n"
            "exp1 107\n"
            "sign1 1\n"
            "sign0 1\n";

    // 初始化配对参数
    pbc_param_t param;
    pbc_param_init_set_str(param, fixed_param_str);
    pairing_t pairing;
    pairing_init_pbc_param(pairing, param);

    // 固定生成元 g
    element_t g;
    element_init_G1(g, pairing);
    const unsigned char g_x[] = {
            0x85, 0x08, 0x24, 0x16, 0x52, 0x13, 0x41, 0x40, 0x62, 0x66, 0x55, 0x16, 0x91, 0x52, 0x49, 0x29,
            0x52, 0x83, 0x12, 0x86, 0x75, 0x74, 0x81, 0x38, 0x50, 0x79, 0x27, 0x54, 0x41, 0x17, 0x06, 0x43,
            0x71, 0x84, 0x65, 0x04, 0x33, 0x83, 0x66, 0x42, 0x63, 0x92, 0x32, 0x85, 0x21, 0x72, 0x80, 0x91,
            0x86, 0x88, 0x89, 0x41, 0x54, 0x32, 0x90, 0x01, 0x95, 0x36, 0x51, 0x66, 0x41, 0x48, 0x79, 0x94,
            0x92, 0x49, 0x11, 0x75, 0x24, 0x77, 0x09, 0x53, 0x42, 0x49, 0x26, 0x36
    };
    const unsigned char g_y[] = {
            0x56, 0x81, 0x08, 0x33, 0x89, 0x02, 0x72, 0x87, 0x31, 0x82, 0x52, 0x20, 0x05, 0x71, 0x25, 0x90,
            0x00, 0x39, 0x86, 0x61, 0x56, 0x15, 0x43, 0x30, 0x62, 0x31, 0x99, 0x89, 0x93, 0x18, 0x56, 0x56,
            0x59, 0x26, 0x40, 0x48, 0x59, 0x35, 0x54, 0x04, 0x07, 0x11, 0x29, 0x79, 0x83, 0x98, 0x84, 0x95,
            0x75, 0x59, 0x22, 0x70, 0x06, 0x33, 0x13, 0x38, 0x17, 0x54, 0x22, 0x78, 0x51, 0x38, 0x05, 0x97,
            0x75, 0x85, 0x99, 0x35, 0x39, 0x04, 0x05, 0x02, 0x47, 0x54, 0x25, 0x91, 0x36
    };
    element_from_bytes(g, (unsigned char*)g_x);
    element_from_bytes_compressed(g, (unsigned char*)g_y);

    // 从控制台输入签名人数 n
    int n;
    std::cout << "Enter the number of signers (n): ";
    std::cin >> n;

    // 生成 n 对公私钥
    std::vector<element_t> sks(n);
    std::vector<element_t> pks(n);
    for (int i = 0; i < n; ++i) {
        init_bls(pairing, g, sks[i], pks[i]);
    }

    // 生成 n 个签名信息
    std::vector<const char*> messages(n);
    for (int i = 0; i < n; ++i) {
        std::string message = "message" + std::to_string(i + 1);
        messages[i] = strdup(message.c_str());
    }

    // 生成 n 个签名
    std::vector<BLSSignature> signatures(n);
    for (int i = 0; i < n; ++i) {
        sign_bls(pairing, g, sks[i], messages[i], signatures[i]);
    }

    // 聚合签名
    BLSSignature aggregated_signature;
    aggregate_signatures(pairing, signatures, aggregated_signature);

    // 序列化公钥并拼接成字符串
    std::vector<std::string> pk_strings;
    for ( auto& pk : pks) {
        pk_strings.push_back(element_to_string(pk));
    }
    std::string pk_string_combined;
    for (size_t i = 0; i < pk_strings.size(); ++i) {
        pk_string_combined += pk_strings[i];
        if (i < pk_strings.size() - 1) {
            pk_string_combined += ",";
        }
    }

    // 序列化聚合签名
    std::string aggregated_signature_string = signature_to_string(aggregated_signature);

    // 输出所有输入数据
    std::cout << "=== Input Data for verify_aggregated_signature ===" << std::endl;
//    std::cout << "Generator g:" << std::endl;
//    element_printf("g = %B\n", g);
    std::cout << "Public Keys (combined): " << pk_string_combined << std::endl;
 //   std::cout << "Messages:" << std::endl;
//    for (size_t i = 0; i < messages.size(); ++i) {
//        std::cout << "  Message" << i + 1 << ": " << messages[i] << std::endl;
//    }
    std::cout << "Aggregated Signature: " << aggregated_signature_string << std::endl;
    std::cout << "=================================================" << std::endl;

    // 验证聚合签名
    bool is_valid = verify_aggregated_signature(pairing, g, pk_string_combined, messages, aggregated_signature_string);
    if (is_valid) {
        std::cout << "Aggregated signature is valid!" << std::endl;
    } else {
        std::cout << "Aggregated signature is invalid!" << std::endl;
    }

//    // test 将公钥第一个字符改为6看签名验证是否能通过
//    pk_string_combined[0] = '6';
//    std::cout << "Modified pk: " << pk_string_combined << std::endl;
//    // 验证聚合签名
//    is_valid = verify_aggregated_signature(pairing, g, pk_string_combined, messages, aggregated_signature_string);
//    if (is_valid) {
//        std::cout << "Aggregated signature is valid!" << std::endl;
//    } else {
//        std::cout << "Aggregated signature is invalid!" << std::endl;
//    }

    // 清理资源
    for (int i = 0; i < n; ++i) {
        element_clear(sks[i]);
        element_clear(pks[i]);
        element_clear(signatures[i].sig);
        free((void*)messages[i]);
    }
    element_clear(g);
    element_clear(aggregated_signature.sig);
    pairing_clear(pairing);

    return 0;
}
