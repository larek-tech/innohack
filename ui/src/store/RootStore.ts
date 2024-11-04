import { ResponseDto } from "@/api/models";

import { makeAutoObservable } from "mobx";

export class RootStore {



    // Stores messages by session ID
    sessionMessages: Map<number, ResponseDto[]> = new Map();

    constructor() {
        makeAutoObservable(this);
        // this.websocketBaseUrl = websocketBaseUrl;
    }

}
