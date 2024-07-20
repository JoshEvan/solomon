import { getBaseUrl, FINDCOMP_URL } from "../../configs/api";
import  Axios from  'axios-observable';
import { Observable } from "rxjs";
import { IIndexItemRequest, IIndexItemResponse, IDeleteItemResponse, IUpsertItemRequest, IUpdateItemRequest } from "../interfaces";

const usingBaseUrl = getBaseUrl()

const serviceIndexItem = (dataPayload:IIndexItemRequest): Observable<IIndexItemResponse> => {
    return Axios.post(
        usingBaseUrl+FINDCOMP_URL.ITEM.INDEX,
        dataPayload
    )
}

export const serviceDeleteItem = (id:string, username:string): Observable<any> => {
    return Axios.delete(
        usingBaseUrl+FINDCOMP_URL.ITEM.DELETE+id+"/"+username
    )
}

export const serviceBuyItem = (id:string, username:string): Observable<any> => {
    return Axios.delete(
        usingBaseUrl+FINDCOMP_URL.ITEM.BUY+id+"/"+username
    )
}

export const serviceAddItem = (dataPayload:IUpsertItemRequest): Observable<any> => {
    return Axios.post(
        usingBaseUrl+FINDCOMP_URL.ITEM.ADD,
        dataPayload
    )
}

export const serviceEditItem = (dataPayload:IUpdateItemRequest): Observable<any> => {
    return Axios.put(
        usingBaseUrl+FINDCOMP_URL.ITEM.EDIT,
        dataPayload 
    )
}

export {serviceIndexItem};