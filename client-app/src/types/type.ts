export interface IResult {
    type: string | null;
    message: string | null;
    iserror: boolean;
    statuscode: number;
    data: IUser[];
  }
export interface IUser {
    userid : string | null;
    username : string | null;
    password : string | null;
    passwordhash : string | null;
    mail : string | null;
    isdeleted: boolean | null;
  }