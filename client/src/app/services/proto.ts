export class ReplyProto { 
    status?: number;     //状态 0正常，小于0出错，大于0可能有问题
    msg?: string;        //状态信息
    data?: any;
    API?: string;        //api接口
    method?: string;     //post,put,get,delete
    // SN?: number;
    rowCount?: number;   //Data若是数组，算其长度
}

export class ReqProto {
    action?: string;     //请求类型GET/POST/PUT/DELETE
    data?: any;          //请求数据
    sets?: string[];
    orderBy?: string;    //排序要求
    filter?: string;     //筛选条件
    page?: number;       //分页
    pageSize?: number;   //分页大小
}