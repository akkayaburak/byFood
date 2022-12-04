export interface IResult {
    type: string | null;
    message: string | null;
    iserror: boolean;
    statuscode: number;
    data: IUser[];
}
export interface IUser {
    userid : string;
    username : string | null;
    password : string | null;
    passwordhash : string | null;
    mail : string | null;
    isdeleted: boolean | null;
}

export interface IUserUpdate {
  userid : string;
  username : string | null;
  mail : string | null;
}

export interface IUserUpdateResponse {
    type: string | null;
    message: string;
    iserror: boolean;
    statuscode: number;
    data: string | null;
}

export interface IUserAdd {
  username : string | null;
  mail : string | null;
  password : string | null;
}
export interface IUserAddResponse {
  type: string | null;
  message: string;
  iserror: boolean;
  statuscode: number;
  data: string | null;
}
