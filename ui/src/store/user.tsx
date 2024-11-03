import { makeAutoObservable, makeObservable } from "mobx";

export interface User {
    Email: string;
    LoggedIn: boolean;
    AccessToken: string;
}


// class RootStore {
//     user: User = {

//     }
//     constructor() {
//         makeAutoObservable(this);
//     }
// }