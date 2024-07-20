import { HTTPCallStatus } from "..";

export interface IInsertItemRequest{
	id:string,
	name:string,
	description:string,
	price:number,
	owner:string,
	category:string
}

export interface IUpdateItemRequest extends IInsertItemRequest{
	itemId:string,
	requester:string
}

export interface IUpsertItemResponse{
	status: HTTPCallStatus,
	data:string[]
}