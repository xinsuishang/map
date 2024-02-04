namespace go gpt

include "../base/common.thrift"

struct ApplicationReq {
    1: i32 modelId (vt.ge = "1")
    2: string sessionId
    3: string prompt (vt.min_size = "1", vt.max_size = "2000")
}

struct ApplicationResp {
    1: i32 modelId
    2: string sessionId
    3: string requestId
    4: string text
}

struct ApplicationListReq {
    1: i32 modelId
    2: i32 parentId
    3: common.Page page
}

struct ApplicationListResp {
    1: list<ApplicationInfo> list
    2: common.Page page
}

struct ApplicationInfo {
    1: i32 modelId
    2: i32 parentId
    3: string name
    4: string model
    5: bool isApplication
    6: string dashboard
    7: string desc
}

service ChatService {
    ApplicationResp Chat(1: ApplicationReq req) ( api.post = '/gpt/application/chat', api.param = 'true')
    ApplicationListResp ChatApplicationList(1: ApplicationListReq req) ( api.post = '/gpt/application/list', api.param = 'true')
}