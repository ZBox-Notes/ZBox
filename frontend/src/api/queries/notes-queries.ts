import { Note } from "@/types/types";
import { useQuery } from "@tanstack/react-query";
import axios from "axios";



export const getAllNotes = (): Promise<Note[]> => {
    return axios.get(import.meta.env.VITE_BACKEND_API_URL + "/notes").then((response) => {
        let notes: Note[] = response.data;
        return notes
    })
}

export const getNote = (id: number): Promise<Note> => {
    return axios.get(import.meta.env.VITE_BACKEND_API_URL + "/notes" + `/${id}`).then((response) => {
        let note: Note = response.data;
        return note
    })
}

export const useAllNotes = () =>
    useQuery({
        queryKey: ["all_notes"],
        queryFn: getAllNotes
    })

export const useNote = (id: number) =>
    useQuery({
        queryKey: [`note_${id}`],
        queryFn: () => getNote(id)
    })