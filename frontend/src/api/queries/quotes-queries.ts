import { Quote } from "@/types/types";
import { useQuery } from "@tanstack/react-query";
import axios from "axios";

const QUOTE_API_URL = "https://zenquotes.io/api/today";

export const getQOTD = (): Promise<Quote> => {
    return axios.get(QUOTE_API_URL).then((response) => {
        let quote: Quote[] = response.data;
        return quote[0]
    })
}

export const useQOTD = () =>
    useQuery({
        queryKey: ["qotd"],
        queryFn: getQOTD
    })