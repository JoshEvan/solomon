import { IItem } from "./IItem";

export interface IIndexItemRequest{
    owner:string,
    category:string
}

export interface IIndexItemResponse{
    data:{items:IItem[]}
}