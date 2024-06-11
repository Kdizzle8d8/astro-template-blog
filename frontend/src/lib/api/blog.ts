import { envField } from "astro/config";

export type Post = {
    id: string;
    createdAt: string;
    updatedAt: string;
    title: string;
    published: boolean;
    author: string;
    description: string;
    content: string;
};

const getBlog = async ()=>{
    return fetch(import.meta.env.VITE_BASE_URL,{
        method:'GET',
        headers:{
            'Content-Type':'application/json',
        }
    }).then((res)=>res.json());
}

getBlog().then((res)=>console.log(res))