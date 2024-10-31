import logo2 from './assets/images/SSLogs.png'

import React, { useEffect, useState } from 'react';
import './App.css';
import { GetDate } from '../wailsjs/go/handlers/Home';
import MessageStats from './components/MessageStats/MessageStats';
import LoadLogs from './components/LoadLogs/LoadLogs';
import styles from './pageCSS/Home.module.css'

function Home() {
    const [date, setDate] = useState('');

    useEffect(() => {
        async function fetchData() {
            try {
                const dateResult = await GetDate();
                setDate(dateResult);
            } catch (err) {
                console.error('Error fetching data:', err);
            }
        }
        fetchData();

        const interval = setInterval(fetchData, 300); // Fetch data every 3 seconds

        return () => clearInterval(interval); // Cleanup interval on component unmount
    }, []);

    return (
        <div className="App">
            <div>
                <LoadLogs />
            </div>

            <div className={styles.statCont}>
                <MessageStats />
            </div>

            <div className={styles.centered}>
                <div className={styles.logoCont}>
                    <img src={logo2} className={styles.logoImg} id="logo" alt="logo"/>
                    <div className={styles.logoTitle}>Super Site Logs</div>
                </div>
            </div>

            <header className="App-header">
                <p>{date}</p>
            </header>
        </div>
    );
}

export default Home;