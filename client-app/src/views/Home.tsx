import { useEffect, useState } from 'react';
import { Space, Table } from 'antd';
import { getUsers } from '../services/service';
import { IUser } from '../types/type';


const { Column } = Table;

const Home = () => {
  const [users, setUsers] = useState<IUser[]>()
  useEffect(() => {
    getUsers(setUsers)
  }, [])
  
 return(
  <>
  <Table dataSource={users} key="dsgs">
    <Column title="Username" dataIndex="username" key="username" />
    <Column title="Mail" dataIndex="mail" key="mail" />
    <Column title="PasswordHash" dataIndex="passwordhash" key="passwordhash" />
    <Column
      title="Action"
      key="action"
      render={(_: any, record: IUser) => (
        <Space size="middle">
          <a>Update </a>
          <a>Delete</a>
        </Space>
      )}
    />
  </Table>
  </>
 );
  
};


export default Home;
