/* Do not change, this code is generated from Golang structs */


export interface ProcessData {
    pid: number;
    title?: string;
    mainModule?: string;
}
export interface ErrorData {
    code?: string;
    message: string;
    stack?: string;
}
export interface ResponseData {
    statusCode: number;
    timeStart: number;
    timeEnd: number;
    headers: HeaderItem[];
    body?: string;
}
export interface HeaderItem {
    name: string;
    value: string;
}
export interface RequestData {
    id: string;
    timeStart: number;
    timeEnd: number;
    url: string;
    method: string;
    remoteAddress?: string;
    headers: HeaderItem[];
    body?: string;
}
export interface LoggerEvent {
    request: RequestData;
    response?: ResponseData;
    error?: ErrorData;
    process?: ProcessData;
}