import { getQOTD } from '@/api/queries/quotes-queries';
import { useQuery } from '@tanstack/react-query';
import React from 'react';
import { Skeleton } from './ui/skeleton';

const Quote: React.FC = () => {
    let quote: string = "If it is doable, I will do it.";
    let author: string = "David René";

    const { data, isLoading, isError } = useQuery({
        queryKey: ["qotd"],
        queryFn: getQOTD,
        staleTime: 1e3 * 60 * 5
    });

    if (!isError) {
        quote = data?.quote?.content || quote;
        author = data?.quote?.author.name || author;
    }

    return (
        <div className='px-8 py-8 bg-gray-100 rounded-(--radius-sm) max-w-2xl'>
            <span className='text-xl text-center italic text-gray-500'>
                <p>
                    {isLoading ?
                        <div><Skeleton className='m-2 w-[600px] h-6 bg-gray-300 mx-auto' />
                            <Skeleton className=' m-2 w-[500px] h-6 bg-gray-300 mx-auto' /></div> :
                        "\"" + quote + "\""
                    }
                </p>
                <p className='mt-2 text-right text-lg text-gray-500'>
                    {isLoading ?
                        <Skeleton className='ml-auto my-2 w-[100px] h-6 bg-gray-300' /> :
                        "- " + data?.quote?.author.name || "No author available."
                    }
                </p>
            </span>
        </div>
    );
};

export default Quote;