export class ReplyProto {
    status?: number;
    msg?: string;
    data?: any;
    API?: string;
    method?: string;
    // SN?: number;
    rowCount?: number;
}

export class ReqProto {
    action?: string;
    data?: any;
    sets?: string[];
    orderBy?: string;
    filter?: string;
    page?: number;
    pageSize?: number;
}