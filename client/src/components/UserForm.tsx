import { BASE_URL } from "../App";
import { Button, Flex, Input, Spinner } from "@chakra-ui/react";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { useState } from "react";
import { IoMdAdd } from "react-icons/io";

const UserForm = () => {
    const [name, setName] = useState("");
    const [email, setEmail] = useState("");

    const queryClient = useQueryClient();

    const { mutate: createUser, isPending: isCreating } = useMutation({
        mutationKey: ["createUser"],
        mutationFn: async (e: React.FormEvent) => {
            e.preventDefault();
            if (!name || !email) return alert("Заполните все поля!");

            try {
                const res = await fetch(BASE_URL + `/users`, {
                    method: "POST",
                    headers: { "Content-Type": "application/json" },
                    body: JSON.stringify({ name, email }),
                });

                const data = await res.json();

                if (!res.ok) {
                    throw new Error(data.error || "Something went wrong");
                }

                setName("");
                setEmail("");
                return data;
            } catch (error: any) {
                throw new Error(error);
            }
        },
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: ["users"] });
        },
        onError: (error: any) => {
            alert(error.message);
        },
    });

    return (
        <form onSubmit={createUser}>
            <Flex gap={2} direction="column">
                <Input
                    type="text"
                    placeholder="Имя"
                    value={name}
                    onChange={(e) => setName(e.target.value)}
                />
                <Input
                    type="text"
                    placeholder="Email"
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                />
                <Button type="submit" mx={2} _active={{ transform: "scale(.97)" }}>
                    {isCreating ? <Spinner size={"xs"} /> : <IoMdAdd size={30} />}
                </Button>
            </Flex>
        </form>
    );
};

export default UserForm;