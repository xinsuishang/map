namespace go oss

include "../base/common.thrift"

struct WeightUploadRequest {
    1: double weight (vt.gt = "0.0", vt.lt = "200.0")
    2: string fileName (vt.min_size = "5", vt.max_size = "100", vt.suffix=".JPG")
    3: string remoteName (vt.min_size = "5", vt.max_size = "100", vt.suffix=".JPG")
    4: bool forceUpload
    5: string dateTime (vt.min_size = "10", vt.max_size = "10")
}

struct NotionUploadRequest {
    1: i32 domainId (vt.gt = "0")
    2: string dateTime (vt.min_size = "10", vt.max_size = "10")
    3: double weight (vt.gt = "0.0", vt.lt = "200.0")
    4: string fileUrl (vt.min_size = "1", vt.max_size = "100")
}

struct UploadRequest {
    1: i32 domainId (vt.gt = "0")
    2: string fileName (vt.min_size = "5", vt.max_size = "100", vt.suffix=".JPG")
    3: string remoteName (vt.min_size = "5", vt.max_size = "100", vt.suffix=".JPG")
    4: string remoteDir (vt.min_size = "1", vt.max_size = "300")
    5: bool forceUpload
}

struct OssUploadResp {
    1: string remoteUrl
}

service UploadService {
    OssUploadResp OssUpload(1: UploadRequest req) ( api.post = '/oss/upload', api.param = 'true')
    common.EmptyResponse NotionUpload(1: NotionUploadRequest req) ( api.post = '/oss/notionUpdate', api.param = 'true')
    common.EmptyResponse WeightUpload(1: WeightUploadRequest req) ( api.post = '/oss/weight', api.param = 'true')
}
