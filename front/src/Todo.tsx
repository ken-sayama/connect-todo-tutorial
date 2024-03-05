import {FC} from "react";
import {useQuery} from "@tanstack/react-query";
import { getTodo } from "./gen/todo/v1/todo-TodoService_connectquery"

type Props = {
    id: string
}

export const GetTodo: FC<Props> = ({ id }) => {
    const { isLoading, isError, error, data } = useQuery(
        getTodo.useQuery({id})
    );

    return (
        <>
            {isLoading && <p>読み込み中</p>}
            {isError && <p>{error?.message}</p>}
            {!isLoading && !isError && data && <h2>{data?.todo?.name}</h2>}
        </>
    )
}