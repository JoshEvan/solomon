
const API_URL = {
    DEV: {
        BASE_URL: 'http://localhost:8080/api/v1/findcomp'
    },
    PROD: {
        BASE_URL:'https://findcomp-josh-be.herokuapp.com/api/v1/findcomp'
    }
}

export const FINDCOMP_URL={
    ITEM: {
        INDEX: '/item/',
        DELETE:'/item/delete/',
        ADD:'/item/insert/',
        EDIT:'/item/update',
        BUY:'/item/buy/',
    },
    CAT:{
        INDEX:'/category/index'
    },
    USER:{
        REGISTER:'/user/register',
        UPDATE:'/user/update',
        SHOW:'/user/show/'
    }
}

export const getBaseUrl = () =>{
    let URL: string;
    if(process.env.NODE_ENV === 'production'){
        URL = API_URL.PROD.BASE_URL;
    }else URL = API_URL.DEV.BASE_URL;
    return URL;
}

export const getLoginUrl = () => {
    if(process.env.NODE_ENV === 'production'){
        return 'https://findcomp-josh-be.herokuapp.com/login'
    }
    return 'http://localhost:8080/login'
}
