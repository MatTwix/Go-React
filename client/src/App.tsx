import { Container, Stack } from '@chakra-ui/react'
import UserForm from './components/UserForm';
import UserList from './components/UserList';

function App() {
  return (
    <Stack h="100vh">
      <Container>
        <UserForm/>
        <UserList/>
      </Container>
    </Stack>
  );
}

export default App
