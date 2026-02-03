/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package edu.eci.arsw.blacklistvalidator;

import java.util.List;

/**
 *
 * @author Daniel palacios Moreno & Sofia Nicolle Ariza Goenaga
 */
public class Main {
    
    public static void main(String[] args){
        Runtime runtime = Runtime.getRuntime();
        long startTime = System.nanoTime();
        
        HostBlackListsValidator hblv=new HostBlackListsValidator();
        int n = 204;
        List<Integer> blackListOcurrences=hblv.checkHost("202.24.34.55", n);
        System.out.println("The host was found in the following blacklists:"+blackListOcurrences);

        long endTime = System.nanoTime();
        long executionTime = (endTime - startTime) / 1_000_000;

        System.out.println("Execution time: " + executionTime + " milliseconds");
    }
    
}
