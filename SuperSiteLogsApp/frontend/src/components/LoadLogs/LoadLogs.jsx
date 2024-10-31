import React, { useEffect, useState } from "react";
import styles from './LoadLogs.module.css'
import { SyncLogs, IsCurrentlyLoading } from '../../../wailsjs/go/handlers/SyncLogs'

function LoadLogs() {
    const [isLoading, setIsLoading] = useState(false)

    async function LoadLogsButtonPress() {
        setIsLoading(true)
        let logsResult = await SyncLogs();
        setIsLoading(false)
    }

    function IsCurrentlyLoadingClass() {
        return isLoading ? styles.currentlyLoading : '';
    }

    async function CheckIfLoading() {
        let isLoading = await IsCurrentlyLoading();
        setIsLoading(isLoading);
        if (isLoading) {
            setTimeout(CheckIfLoading, 500);
        }
    }

    useEffect(() => {
        CheckIfLoading();
    }, [])

    return (
        <div className={styles.loadLogsCont}>
            <div className={styles.loadLogsButton}>
                <button className={`${IsCurrentlyLoadingClass()} ${ styles.loadLogsIcon} fas fa-sync`} onClick={LoadLogsButtonPress}></button>
            </div>
        </div>
    )
}

export default LoadLogs