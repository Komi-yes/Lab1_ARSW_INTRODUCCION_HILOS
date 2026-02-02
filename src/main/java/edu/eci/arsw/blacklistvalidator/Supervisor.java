package edu.eci.arsw.blacklistvalidator;

import edu.eci.arsw.spamkeywordsdatasource.HostBlacklistsDataSourceFacade;

import static edu.eci.arsw.blacklistvalidator.HostBlackListsValidator.*;

/*
 *
 * @author Daniel palacios Moreno & Sofia Nicolle Ariza Goenaga
 */
public class Supervisor extends Thread{
    private String host;
    private int section;
    private int sectionSize;

    public Supervisor(String host, int section, int sectionSize){
        this.host = host;
        this.section = section;
        this.sectionSize = sectionSize;
    }

    public void run(){
        try {
            HostBlacklistsDataSourceFacade skds = HostBlacklistsDataSourceFacade.getInstance();

            for (int i = section * sectionSize; i < (section + 1) * sectionSize; i++){
                if (!shouldContinueSearching()) {
                    break;
                }

                incrementCheckedCount();
                if (skds.isInBlackListServer(i, host)){
                    addBlackListOcurrence(i);
                }
            }
        } finally {
            notifyThreadCompletion();
        }
    }
}