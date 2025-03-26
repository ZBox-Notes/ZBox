import { Box } from "@/types/types";
import { useQuery } from "@tanstack/react-query";
import axios from "axios";



export const getAllBoxes = (): Promise<Box[]> => {
    return axios.get(import.meta.env.VITE_BACKEND_API_URL + "/boxes").then((response) => {
        let boxes: Box[] = response.data;
        return boxes
    })
}

export const getBox = (id: number): Promise<Box> => {
    return axios.get(import.meta.env.VITE_BACKEND_API_URL + "/boxes" + `/${id}`).then((response) => {
        let box: Box = response.data;
        return box
    })
}

export const useAllBoxes = () =>
    useQuery({
        queryKey: ["all_boxes"],
        queryFn: getAllBoxes
    })

export const useBox = (id: number) =>
    useQuery({
        queryKey: [`box_${id}`],
        queryFn: () => getBox(id)
    })