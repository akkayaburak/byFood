import { useEffect, useState } from 'react';
import { Row, Col, Button, Space, Table, Form, Input, Modal } from 'antd';
import { deleteUser, getUsers, updateUser, createUser } from '../services/service';
import { IUser, IUserUpdate, IUserAdd } from '../types/type';


const { Column } = Table;

const Home = () => {
  const [addUserForm] = Form.useForm();
  const [users, setUsers] = useState<IUser[]>();
  const [selectedUser, setSelectedUser] = useState<IUser>();
  const [selectedId, setSelectedId] = useState<string>();
  const [isDeleteModalOpen, setIsDeleteModalOpen] = useState(false);
  const [isUpdateModalOpen, setIsUpdateModalOpen] = useState(false);
  const [respMessage, setRespMessage] = useState<string>();

  const [addUserModal, setAddUserModal] = useState<boolean>(false);

  useEffect(() => {
    getUsers(setUsers)
  }, []);

  //onFinish func. for add new user start
  const onFormFinishAddUser = (values: any) => {
    const body: IUserAdd = {
      username: values.username,
      password: values.password,
      mail: values.mail,
    }
    createUser(body, setAddUserModal, info, setUsers);
    addUserForm.resetFields();
  };
  //onFinish func. for add new user end

  //onFinish func. for updated selected user start
  const onFormFinishUpdate = (values: any) => {
    const updateUserObj: IUserUpdate = {
      userid: selectedUser?.userid || '',
      username: values.username,
      mail: values.mail,
    }
    updateUser(setUsers, updateUserObj, setIsUpdateModalOpen, setRespMessage);
  };
  const onFinishFailed = (errorInfo: any) => {
    console.log('Failed:', errorInfo);
  };
  //onFinish func. for updated selected user start

  //delete user functions start
  const showDeleteModal = (userid : string) => {
    setIsDeleteModalOpen(true);
    setSelectedId(userid)
  };
  const handleDeleteOk = () => {
    setIsDeleteModalOpen(false);
    deleteUser(selectedId || null)
    setUsers(users => users?.filter(user => user.userid !== selectedId));
  };
  const handleDeleteCancel = () => {
    setIsDeleteModalOpen(false);
  };
  //delete user functions end

  //update user functions end
  const showUpdateModal = (user: IUser) => {
    setIsUpdateModalOpen(true);
    setSelectedUser(user)
  };
  const handleUpdateOk = () => {
    setIsUpdateModalOpen(false);
    setUsers(users => users?.filter(user => user.userid !== selectedId));
  };
  const handleUpdateCancel = () => {
    setIsUpdateModalOpen(false);
  };
  const handleNewPerson = () => {
    setAddUserModal((prev) => !prev);
  }
  //update user functions end

  // create info modal for success add func. start
  const info = () => {
    Modal.info({
      content: (
        <div>
          <p>The user has been inserted successfully!</p>
        </div>
      ),
      onOk() {},
    });
  };
  // create info modal for success add func. end

 return(
  <>
    <Button type="primary" onClick={() => handleNewPerson()}>+ Add User</Button>
    <Table dataSource={users} key="table" pagination={false}>
      <Column title="Username" dataIndex="username" key="username" />
      <Column title="Mail" dataIndex="mail" key="mail" />
      <Column title="PasswordHash" dataIndex="passwordhash" key="passwordhash" />
      <Column
        title="Action"
        key="action"
        render={(_: any, record: IUser) => (
          <Space size="middle">
            <Button 
              onClick={() => showUpdateModal(record)}
            >
              Update
            </Button>
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
      cancelText="Back"
    >
      <p>Are you sure you want to delete this user?</p>
    </Modal>
    {/* update user modal start */}
    <Modal
      open={isUpdateModalOpen}
      onCancel={handleUpdateCancel}
      footer={null}
      >
        {respMessage ? 
          (<div><p>{respMessage}</p></div>)
        :
          (
            <>
              <div>
                <p>You are updating this user : {selectedUser?.username}</p>
              </div>
              <Row>
                <Col span={24} style={{ textAlign: 'right' }}>
                  <Form
                    name="updateForm"
                    labelCol={{ span: 8 }}
                    wrapperCol={{ span: 16 }}
                    onFinish={onFormFinishUpdate}
                    onFinishFailed={onFinishFailed}
                    autoComplete="off"
                  >
                    
                    <Form.Item
                      label="Username"
                      name="username"
                    >
                      <Input />
                    </Form.Item>

                    <Form.Item
                      label="Mail"
                      name="mail"
                    >
                      <Input />
                    </Form.Item>
                    <Form.Item wrapperCol={{ offset: 8, span: 16 }}>
                      <Button type="primary" htmlType="submit">
                        Update
                      </Button>
                    </Form.Item>
                  </Form>
                </Col>  
              </Row>
            </>
          )
        }
    </Modal>
    {/* update user modal end */}

    {/* add user modal start */}
    <Modal 
      title="Add User"
      open={addUserModal} 
      onCancel={() => setAddUserModal((prev) => !prev)}
      footer={null}
    >
      <Row>
        <Col span={24} style={{ textAlign: 'right' }}>
          <Form
            name="addUserForm"
            labelCol={{ span: 8 }}
            wrapperCol={{ span: 16 }}
            onFinish={onFormFinishAddUser}
            onFinishFailed={onFinishFailed}
            autoComplete="off"
          >
            <Form.Item
              label="Username"
              name="username"
              rules={[{ required: true, message: 'Please input your username!' }]}
            >
              <Input />
            </Form.Item>
            <Form.Item
              label="Password"
              name="password"
              rules={[{ required: true, message: 'Please input your password!' }]}
            >
              <Input />
            </Form.Item>
            <Form.Item
              label="Mail"
              name="mail"
              rules={[{ required: true, message: 'Please input your mail!' }]}
            >
              <Input />
            </Form.Item>
            <Form.Item wrapperCol={{ offset: 8, span: 16 }}>
              <Button type="primary" htmlType="submit">
                Add
              </Button>
            </Form.Item>
          </Form>
        </Col>  
      </Row>
    </Modal>
    {/* add user modal end */}
  </>
 );
};


export default Home;
