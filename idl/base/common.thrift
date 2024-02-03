namespace go common

struct EmptyResponse {

}

struct Page {
    1: i32 pageNo (vt.gt = "1")
    2: i32 pageSize (vt.gt = "1", vt.lt = "100")
    3: i32 total
}