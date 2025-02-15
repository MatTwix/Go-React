import { Flex, Spinner, Stack, Text } from "@chakra-ui/react";
import UserListItem from "./UserListItem";
import { useQuery } from "@tanstack/react-query";
import { BASE_URL } from "../App";

export type User = {
    id: number;
    name: string;
    email: string;
}

const UserList = () => {
	const {data:users, isLoading} = useQuery<User[]>({
        queryKey:["users"],
        queryFn: async () => {
            try {
                const res = await fetch(BASE_URL + `/users`)
                const data = await res.json()

                if (!res.ok) {
                    throw new Error(data.error || "Something went wrong")
                }
                return data || []
                } catch (error) {
                    console.log(error)
                }
            }
    })

	return (
		<>
			<Text fontSize={"4xl"} textTransform={"uppercase"} fontWeight={"bold"} textAlign={"center"} my={2}>
				Available users
			</Text>
			{isLoading && (
				<Flex justifyContent={"center"} my={4}>
					<Spinner size={"xl"} />
				</Flex>
			)}
			{!isLoading && users?.length === 0 && (
				<Stack alignItems={"center"} gap='3'>
					<Text fontSize={"xl"} textAlign={"center"} color={"gray.500"}>
                        There are no users yet!
					</Text>
				</Stack>
			)}
			<Stack gap={3}>
				{users?.map((user) => (
					<UserListItem key={user.id} user={user} />
				))}
			</Stack>
		</>
	);
};
export default UserList;