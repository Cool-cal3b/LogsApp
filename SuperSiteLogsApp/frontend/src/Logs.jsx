import React, { useEffect, useState } from 'react';
import { GetLogs, SetShowByType, GetShowByType, MarkAsRead, SetLogLevel, GetLogLevel, MarkAllAsRead } from '../wailsjs/go/handlers/Logs'; 
import styles from './pageCSS/Logs.module.css'
import { Link } from 'react-router-dom';

function Logs() {
    const [logs, setLogs] = useState([]);
    const [page, setPage] = useState(0);
    const [showBy, setShowBy] = useState(0);
    const [logLevel, setLogLevel] = useState(0);
    const [totalLogs, setTotalLogs] = useState(0)
    const [isLastPage, setIsLastPage] = useState(false)
    const [syncLogs, doSyncLogs] = useState(0);

    useEffect(() => {
        async function fetchLogs() {
            try {
                const logsResult = await GetLogs(page);
                setLogs(logsResult.Logs);
                setTotalLogs(logsResult.TotalCount)
                setIsLastPage(logsResult.IsLastPage)
            } catch (err) {
                console.error('Error fetching logs:', err);
            }
        }
        fetchLogs();

        async function getShowBy() {
            let showByFromServer = await GetShowByType();
            setShowBy(showByFromServer);
        }
        getShowBy();

        async function getLogLevelFromServer() {
            let logLevelFromServer = await GetLogLevel();
            setLogLevel(logLevelFromServer);
        }
        getLogLevelFromServer();
    }, [page, showBy, logLevel, syncLogs]);

    function getLogErrorClass(log) {
        return log.IsError ? styles.logErrorMessage : '';
    }

    function getUnviewedMessageClass(log) {
        return !log.Read ? styles.unviewedMessage : '';
    }

    function leftArrow() {
        if (page <= 0) {
            return;
        }

        setPage(p => p-1)
    }

    function rightArrow() {
        if (isLastPage) {
            return;
        }

        setPage(p => p+1)
    }

    function formatDate(date) {
        return new Date(date).toLocaleString('en-US', {
            month: 'short',
            day: 'numeric',
            year: 'numeric',
            hour: 'numeric',
            minute: 'numeric',
            hour12: true
          })
    }

    async function SaveShowBy(e) {
        let selectedOption = parseInt(e.target.value);  
        let result = await SetShowByType(selectedOption);
        setShowBy(selectedOption);
        setPage(0)
    }

    async function SaveLogLevel(e) {
        let selectedOption = parseInt(e.target.value);
        let result = await SetLogLevel(selectedOption);
        setLogLevel(selectedOption);
        setPage(0)
    }

    async function setLogAsRead(id) {
        let result = await MarkAsRead(id)
        doSyncLogs(x => x+1)
    }

    async function markAllAsRead() {
        let result = await MarkAllAsRead();
        doSyncLogs(x => x+1)
    }

    return (
        <div className={styles.logsPage}>
            <h1>Logs Page</h1>

            <div className={styles.selectHolder}>
                <select className={styles.logsToShowSelect} onChange={SaveShowBy} value={showBy}>
                    <option value='0'>Show All</option>
                    <option value='1'>Show Non-read</option>
                </select>

                <select className={styles.logsToShowSelect} onChange={SaveLogLevel} value={logLevel}>
                    <option value='3'>Show All</option>
                    <option value='0'>Show Messages</option>
                    <option value='1'>Show Warnings</option>
                    <option value='2'>Show Errors</option>
                </select>
            </div>

            <div className={styles.tableSection}>
                <div className={styles.markAllButtonCont}>
                    <input type='button' value='Mark all as read' className={styles.markAllAsRead} onClick={markAllAsRead} />
                </div>

                <div className={styles.logsTable}>
                    <div className={styles.arrowCont}>
                        <div>Page {page + 1}</div>
                        <div className={`${styles.logArrows} fas fa-arrow-left`} onClick={leftArrow}></div>
                        <div className={`${styles.logArrows} fas fa-arrow-right`} onClick={rightArrow}></div>
                        <div>({totalLogs} Logs)</div>
                    </div>
                    <table>
                        <thead>
                            <tr>
                                <th></th>
                                <th><div className={styles.logHeader}>Date</div></th>
                                <th><div className={styles.logHeader}>Message</div></th>
                                <th></th>
                            </tr>
                        </thead>
                        <tbody>
                            {logs.map((log, index) => (
                                <tr key={index}>
                                    <td><div className={getUnviewedMessageClass(log)}></div></td>
                                    <td>
                                        <div>{formatDate(log.Time)}</div>
                                    </td>
                                    <td>
                                        <Link to={`/log/${log.ID}`} className={`${styles.logMessageLink} ${getLogErrorClass(log)}`}>
                                            <div className={styles.logMessage}>{log.Message}</div>
                                        </Link>
                                    </td>
                                    <td><div className={styles.viewedButton} onClick={() => setLogAsRead(log.ID)}>Viewed</div></td>
                                </tr>
                            ))}
                        </tbody>
                    </table>
                </div>
            </div>
        </div>
    );
}

export default Logs;
