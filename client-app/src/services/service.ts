import axios from "axios"
import { IResult, IUser } from "../types/type"


export const getUsers = (setUsers: (value: IUser[]) => void) => {
    axios.get<IResult>('http://localhost:8080/users/')
      .then((res) => {
        setUsers(res.data.data)
      }).catch((err) =>  console.log(err))
} 

export const deleteUser = (userid: string | null) => {
    axios.delete<IResult>(`http://localhost:8080/users/`,{data: {"userid" : userid}})
      .then((res) => {
      }).catch((err) =>  console.log(err))
} 

