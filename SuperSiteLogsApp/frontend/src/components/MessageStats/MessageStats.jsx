import React, { useEffect, useState } from 'react';
import { GetErrorAndMessageCount } from '../../../wailsjs/go/handlers/LogsUtils'
import styles from './MessageStats.module.css'

const ErrorMessageCount = () => {
    const [errors, setErrors] = useState(0);
    const [messages, setMessages] = useState(0);

    useEffect(() => {
        async function fetchData() {
            try {
                let result = await GetErrorAndMessageCount();

                setErrors(result.errors); 
                setMessages(result.messages); 
            } catch (err) {
                console.error('Error fetching data:', err);
            }
        }
        fetchData();
    }, []);

    return (
        <div className={styles.statCont}>
            <p className={`${styles.statCard} ${styles.redCard}`}>Errors: {errors}</p>
            <p className={`${styles.statCard} ${styles.greenCard}`}>Messages: {messages}</p>
        </div>
    );
};

export default ErrorMessageCount;
