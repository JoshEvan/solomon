export enum HTTPCallStatus{
    Success = 'SUCCESS',
    Failed = 'FAILED'
}
export interface BaseResponse{
    status:HTTPCallStatus,
    data:{}
}

export { ICRUDResponse } from './ICRUD'

export { IItem } from './items/IItem';
export { IIndexItemRequest, IIndexItemResponse } from './items/IIndexItem';
export { IUpdateItemRequest, IUpsertItemResponse, IInsertItemRequest } from './items/IUpsertItem';