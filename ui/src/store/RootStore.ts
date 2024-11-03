import { QueryDto, ResponseDto } from "@/api/models";
import { LOCAL_STORAGE_KEY } from "@/auth/AuthProvider";
import { makeAutoObservable, runInAction } from "mobx";

export class RootStore {



    // Stores messages by session ID
    sessionMessages: Map<number, ResponseDto[]> = new Map();

    constructor() {
        makeAutoObservable(this);
        // this.websocketBaseUrl = websocketBaseUrl;
    }

}
