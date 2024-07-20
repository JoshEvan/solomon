export interface ILoginRequest{
    username: string,
    password: string
}

export interface IRegisterRequest extends ILoginRequest{
    profileInfo: string
}

export const convIRegisterRequestToILoginRequest = (data : IRegisterRequest) => {
    return {
        username:data.username,
        password:data.password
    }
}

export interface IUpdateUserRequest extends IRegisterRequest{
    newPassword:string
}