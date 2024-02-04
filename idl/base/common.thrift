namespace go common

struct EmptyResponse {

}

struct Page {
    1: i32 pageNo (vt.ge = "1")
    2: i32 pageSize (vt.ge = "1", vt.le = "100")
    3: i32 total
}