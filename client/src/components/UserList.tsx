import { Flex, Spinner, Stack, Text } from "@chakra-ui/react";
import { useState } from "react";
import UserListItem from "./UserListItem";

const UserList = () => {
	const [isLoading, setIsLoading] = useState(false);
	const users = [
		{
			_id: 1,
			body: "Buy groceries",
			completed: true,
		},
		{
			_id: 2,
			body: "Walk the dog",
			completed: false,
		},
		{
			_id: 3,
			body: "Do laundry",
			completed: false,
		},
		{
			_id: 4,
			body: "Cook dinner",
			completed: true,
		},
	];
	return (
		<>
			<Text fontSize={"4xl"} textTransform={"uppercase"} fontWeight={"bold"} textAlign={"center"} my={2}>
				Today's Tasks
			</Text>
			{isLoading && (
				<Flex justifyContent={"center"} my={4}>
					<Spinner size={"xl"} />
				</Flex>
			)}
			{!isLoading && users?.length === 0 && (
				<Stack alignItems={"center"} gap='3'>
					<Text fontSize={"xl"} textAlign={"center"} color={"gray.500"}>
						All tasks completed! 🤞
					</Text>
					<img src='/go.png' alt='Go logo' width={70} height={70} />
				</Stack>
			)}
			<Stack gap={3}>
				{users?.map((user) => (
					<UserListItem key={user._id} user={user} />
				))}
			</Stack>
		</>
	);
};
export default UserList;