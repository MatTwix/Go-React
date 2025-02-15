import { BASE_URL } from "../App";
import { Box, Flex, Spinner, Text } from "@chakra-ui/react";
import { QueryClient, useMutation, useQueryClient } from "@tanstack/react-query";
import { FaCheckCircle } from "react-icons/fa";
import { MdDelete } from "react-icons/md";
import { User } from "./UserList";

const UserItem = ({ user }: { user: User }) => {
    const queryClient = useQueryClient();
    const { mutate: updateUser, isPending: isUpdating } = useMutation({
        mutationKey: ["updateUser"],
        mutationFn: async () => {
            try {
                const resUser = await fetch(`${BASE_URL}/users/${user.id}`);
                if (!resUser.ok) throw new Error("Failed to fetch user data");
    
                const currentUser = await resUser.json();
    
                const res = await fetch(`${BASE_URL}/users/${user.id}`, {
                    method: "PUT",
                    headers: { "Content-Type": "application/json" },
                    body: JSON.stringify({
                        ...currentUser,
                        email: "CHANGED EMAIL" + Math.floor(Math.random() * 10),
                    }),
                });
    
                const data = await res.json();
                if (!res.ok) {
                    throw new Error(data.error || "Something went wrong");
                }
                return data;
            } catch (error) {
                console.log(error);
            }
        },
        onSuccess: () => {
            queryClient.invalidateQueries({queryKey:["users"]})
        }
    });

    const { mutate:deleteUser, isPending:isDeleting } = useMutation({
        mutationKey:["deleteUser"],
        mutationFn: async () => {
            try {
                const res = await fetch(`${BASE_URL}/users/${user.id}`, {
                    method:"DELETE"
                });

                const data = await res.json();

                if (!res.ok) {
                    throw new Error(data.error || "Something went wrong");
                }
                return data;
            } catch (error) {
                console.log(error);
            }
        },
        onSuccess: () => {
            queryClient.invalidateQueries({queryKey:["users"]})
        }    
    })

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
                <Text color={"green.200"}>
                    {user.name} - {user.email}
                </Text>
            </Flex>
            <Flex gap={2} alignItems={"center"}>
                <Box color={"green.500"} cursor={"pointer"} onClick={() => updateUser()}>
                    {!isUpdating && <FaCheckCircle size={20} />}
                    {isUpdating && <Spinner size={"sm"} />}
                </Box>
                <Box color={"red.500"} cursor={"pointer"} onClick={() => deleteUser()}>
                    {!isDeleting && <MdDelete size={25} />}
                    {isDeleting && <Spinner size={"sm"} />}
                </Box>
            </Flex>
        </Flex>
    );
};

export default UserItem;