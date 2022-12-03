import { useEffect, useState } from 'react';
import { Button, Space, Table } from 'antd';
import { deleteUser, getUsers } from '../services/service';
import { IUser } from '../types/type';
import Modal from 'antd/es/modal/Modal';


const { Column } = Table;

const Home = () => {
  const [users, setUsers] = useState<IUser[]>()
  useEffect(() => {
    getUsers(setUsers)
  }, [])
  const [selectedData, setSelectedData] = useState<string>();
  const [isDeleteModalOpen, setIsDeleteModalOpen] = useState(false);

  const showDeleteModal = (userid : string) => {
    setIsDeleteModalOpen(true);
    setSelectedData(userid)
  };

  const handleDeleteOk = () => {
    setIsDeleteModalOpen(false);
    deleteUser(selectedData || null)
    setUsers(users => users?.filter(user => user.userid !== selectedData));

  };

  const handleDeleteCancel = () => {
    setIsDeleteModalOpen(false);
  };



  
 return(
  <>
  <Table dataSource={users} key="table" pagination={false}>
    <Column title="Username" dataIndex="username" key="username" />
    <Column title="Mail" dataIndex="mail" key="mail" />
    <Column title="PasswordHash" dataIndex="passwordhash" key="passwordhash" />
    <Column
      title="Action"
      key="action"
      render={(_: any, record: IUser) => (
        <Space size="middle">
          <a>Update </a>
          <Button 
            onClick={() => showDeleteModal(record.userid)}
            danger
          >
            Delete
          </Button>
        </Space>
      )}
    />
  </Table>

  <Modal 
    open={isDeleteModalOpen} 
    onOk={handleDeleteOk} 
    onCancel={handleDeleteCancel} 
    cancelText="Back">
        <p>Are you sure you want to delete this user?</p>
  </Modal>
  </>
 );
  
};


export default Home;
