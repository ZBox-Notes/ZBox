import { QuoteResponse } from "@/types/types";
import axios from "axios";

const QUOTE_API_URL = "https://api.quotable.kurokeita.dev/api/quotes/random";

export const getQOTD = (): Promise<QuoteResponse> => {
    return axios.get(QUOTE_API_URL).then((response) => {
        let quote: QuoteResponse = response.data;
        return quote
    })
}
