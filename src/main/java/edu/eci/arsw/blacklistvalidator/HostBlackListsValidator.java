package edu.eci.arsw.blacklistvalidator;

import edu.eci.arsw.spamkeywordsdatasource.HostBlacklistsDataSourceFacade;
import java.util.LinkedList;
import java.util.List;
import java.util.concurrent.CountDownLatch;
import java.util.concurrent.TimeUnit;
import java.util.concurrent.atomic.AtomicBoolean;
import java.util.concurrent.atomic.AtomicInteger;
import java.util.logging.Level;
import java.util.logging.Logger;

/*
 *
 * @author Daniel palacios Moreno & Sofia Nicolle Ariza Goenaga
 */

public class HostBlackListsValidator {

    private static final int BLACK_LIST_ALARM_COUNT = 5;
    private static LinkedList<Integer> blackListOcurrences;
    private static AtomicInteger ocurrencesCount;
    private static CountDownLatch stopLatch;
    private static CountDownLatch completionLatch;
    private static AtomicBoolean stopFlag;
    private static AtomicInteger checkedListsCount;

    /**
     * Check the given host's IP address in all the available black lists,
     * and report it as NOT Trustworthy when such IP was reported in at least
     * BLACK_LIST_ALARM_COUNT lists, or as Trustworthy in any other case.
     * The search is not exhaustive: When the number of occurrences is equal to
     * BLACK_LIST_ALARM_COUNT, the search is finished, the host reported as
     * NOT Trustworthy, and the list of the five blacklists returned.
     * @param ipaddress suspicious host's IP address.
     * @return  Blacklists numbers where the given host's IP address was found.
     */
    public List<Integer> checkHost(String ipaddress, int n){
        blackListOcurrences = new LinkedList<>();
        ocurrencesCount = new AtomicInteger(0);
        checkedListsCount = new AtomicInteger(0);
        stopFlag = new AtomicBoolean(false);


        stopLatch = new CountDownLatch(BLACK_LIST_ALARM_COUNT);
        completionLatch = new CountDownLatch(n);

        HostBlacklistsDataSourceFacade skds = HostBlacklistsDataSourceFacade.getInstance();
        int threadSectionSize = skds.getRegisteredServersCount() / n;


        for (int i=0; i < n;i++){
            Thread thread = new Supervisor(ipaddress, i,  threadSectionSize);
            thread.start();
        }

        try {
            if (stopLatch.await(0, TimeUnit.MILLISECONDS)) {
                stopFlag.set(true);
            }

            completionLatch.await();

        } catch (InterruptedException e) {
            Thread.currentThread().interrupt();
            throw new RuntimeException(e);
        }

        if (ocurrencesCount.get() >= BLACK_LIST_ALARM_COUNT){
            skds.reportAsNotTrustworthy(ipaddress);
        } else {
            skds.reportAsTrustworthy(ipaddress);
        }

        Logger LOG = Logger.getLogger(HostBlackListsValidator.class.getName());
        LOG.log(Level.INFO, "Checked Black Lists:{0} of {1}", new Object[]{checkedListsCount.get(), skds.getRegisteredServersCount()});

        return blackListOcurrences;
    }

    public static synchronized void addBlackListOcurrence(int index){
        if (ocurrencesCount.get() < BLACK_LIST_ALARM_COUNT) {
            blackListOcurrences.add(index);
            int count = ocurrencesCount.incrementAndGet();

            if (count <= BLACK_LIST_ALARM_COUNT) {
                stopLatch.countDown();
            }

            if (count >= BLACK_LIST_ALARM_COUNT) {
                stopFlag.set(true);
            }
        }
    }

    public static void notifyThreadCompletion(){
        completionLatch.countDown();
    }

    public static void incrementCheckedCount(){
        checkedListsCount.incrementAndGet();
    }

    public static boolean shouldContinueSearching(){
        return !stopFlag.get();
    }
}