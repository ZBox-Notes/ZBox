export type Note = {
    id: number,
    user_id: number,
    title: string,
    content: string,
    keep_in_inbox: boolean,
    created_at: string,
};

export type Box = {
    id: number,
    user_id: number,
    name: string,
    created_at: string,
    updated_at: string
}

export type QuoteResponse = {
    quote: Quote
}

export type Quote = {
    id: string,
    content: string
    author: QuoteAuthor
}

export type QuoteAuthor = {
    name: string
}

