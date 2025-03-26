import {
    Card,
    CardDescription,
    CardHeader,
    CardTitle
} from "@/components/ui/card";
import { Box } from "@/types/types";
import React from "react";

interface BoxPreviewCardProps {
    box: Box
}

export const BoxPreviewCard: React.FC<BoxPreviewCardProps> = (props: { box: Box }): JSX.Element => {
    let box = props.box
    return (
        <Card>
            <CardHeader>
                <CardTitle>
                    {box.name}
                </CardTitle>
                <CardDescription>
                    {box.created_at}
                </CardDescription>
            </CardHeader>
        </Card>
    )
}