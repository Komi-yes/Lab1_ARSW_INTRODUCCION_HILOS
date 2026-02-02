package edu.eci.arsw.blacklistvalidator;

import edu.eci.arsw.spamkeywordsdatasource.HostBlacklistsDataSourceFacade;
import static edu.eci.arsw.blacklistvalidator.HostBlackListsValidator.addBlackListOcurrence;
import static edu.eci.arsw.blacklistvalidator.HostBlackListsValidator.addCheckedListsCount;

public class Supervisor extends Thread{
    private String host;
    private int section;
    private int sectionSize;

    public Supervisor(String Host, int section, int sectionSize){
        this.section = section;
        this.sectionSize = sectionSize;
        this.host = Host;
    }

    public void run(){
        HostBlacklistsDataSourceFacade skds=HostBlacklistsDataSourceFacade.getInstance();

        for (int i=section*sectionSize;i<(section+1)*sectionSize;i++){
            addCheckedListsCount();
            if (skds.isInBlackListServer(i, host)){
                addBlackListOcurrence(i);
            }
        }
    }
}
