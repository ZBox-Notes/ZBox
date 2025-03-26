import { useParams } from "react-router";

export default function Note() {
    let params = useParams();
    return <>
        {params.note_id}
    </>
}