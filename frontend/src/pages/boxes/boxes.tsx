import { useAllBoxes } from '@/api/queries/boxes-queries';
import { BoxPreviewCard } from '@/components/box-preview-card';
import { Archive } from 'lucide-react';
import React from 'react';

const BoxesPage: React.FC = () => {
    const { data, isSuccess, error, isLoading } = useAllBoxes();
    if (error) console.log(error);

    return (
        <div>
            <div className='flex items-center align-center space-x-4 mb-4'>
                <Archive size={48} strokeWidth={1} />
                <h1>Boxes</h1>
            </div>
            {isLoading ? <div>Loading...</div> : (isSuccess ? (
                (data && data.length) > 0 ? (
                    data.map(box => (
                        <BoxPreviewCard key={box.id} box={box} />
                    ))
                ) : (
                    <div>
                        <p className='italic'>
                            No boxes yet
                        </p>
                    </div>
                )
            ) : (
                <div>
                    <p className='italic'>
                        Oops... Unable to fetch boxes
                    </p>
                </div>
            ))}
        </div>
    );
};

export default BoxesPage;
