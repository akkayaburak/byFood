import axios from "axios"
import { 
  IResult, 
  IUser, 
  IUserUpdate, 
  IUserUpdateResponse, 
  IUserAdd, 
  IUserAddResponse 
} from "../types/type"


export const getUsers = (setUsers: (value: IUser[]) => void) => {
    axios.get<IResult>('http://localhost:8080/users/')
      .then((res) => {
        setUsers(res.data.data)
      }).catch((err) =>  console.log(err))
} 

export const createUser = (
    body: IUserAdd,
    setAddUserModal: (value: boolean) => void,
    info: () => void,
    setUsers: (value: IUser[]) => void, 
  ) => {
    axios.post<IUserAddResponse>('http://localhost:8080/users/', body)
      .then((res) => {
        if (res.data.statuscode === 200) {
          setAddUserModal(false);
          info();
          getUsers(setUsers);
        }
      }).catch((err) =>  console.log(err))
} 

export const deleteUser = (userid: string | null) => {
    axios.delete<IResult>(`http://localhost:8080/users/`,
    {
      data: {"userid" : userid}
    }).then((res) => {
      }).catch((err) =>  console.log(err))
} 

export const updateUser = (
  setUsers: (value: IUser[]) => void, 
  updateUserInfo: IUserUpdate, 
  setIsUpdateModalOpen: (value: boolean) => void,
  setRespMessage: (value: string) => void,
) => {
    const body = {
      "userid" : updateUserInfo.userid, 
      "username" : updateUserInfo.username, 
      "mail": updateUserInfo.mail,
    }
    axios.put<IUserUpdateResponse>(`http://localhost:8080/users/`, body)
    .then((res) => {
      if (res.data.statuscode === 200) {
        getUsers(setUsers);
        setRespMessage(res.data.message)
        setTimeout(() => {
          setIsUpdateModalOpen(false);
        }, 3000);
        
      }
    }).catch((err) => console.log(err))
}

