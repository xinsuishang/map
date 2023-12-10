namespace go user

include "../base/common.thrift"
include "../base/user.thrift"

struct LoginRequest {
    1: i32 uid
    2: string password
}

struct UserResponse {
    2: user.UserInfo data
}

struct RegisterRequest {
    1: string user_name
    2: string password
    3: string phone_number
    4: string avatar_url
    5: string email
    6: string personal_signature
}

struct UserInfoRequest {
    1: i32 uid,
    2: string type, // 查询类型
    3: string key // 邮箱、电话号码等
}

service UserService {
    UserResponse Login(1: LoginRequest req)
    UserResponse Register(1: RegisterRequest req)
    UserResponse UserInfo(1: UserInfoRequest req)
}