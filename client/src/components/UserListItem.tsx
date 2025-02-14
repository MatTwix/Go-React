import { Badge, Box, Flex, Text } from "@chakra-ui/react";
import { FaCheckCircle } from "react-icons/fa";
import { MdDelete } from "react-icons/md";

const UserItem = ({ user }: { user: any }) => {
	return (
		<Flex gap={2} alignItems={"center"}>
			<Flex
				flex={1}
				alignItems={"center"}
				border={"1px"}
				borderColor={"gray.600"}
				p={2}
				borderRadius={"lg"}
				justifyContent={"space-between"}
			>
				<Text
					color={user.completed ? "green.200" : "yellow.100"}
					textDecoration={user.completed ? "line-through" : "none"}
				>
					{user.body}
				</Text>
				{user.completed && (
					<Badge ml='1' colorScheme='green'>
						Done
					</Badge>
				)}
				{!user.completed && (
					<Badge ml='1' colorScheme='yellow'>
						In Progress
					</Badge>
				)}
			</Flex>
			<Flex gap={2} alignItems={"center"}>
				<Box color={"green.500"} cursor={"pointer"}>
					<FaCheckCircle size={20} />
				</Box>
				<Box color={"red.500"} cursor={"pointer"}>
					<MdDelete size={25} />
				</Box>
			</Flex>
		</Flex>
	);
};
export default UserItem;