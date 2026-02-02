/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package edu.eci.arsw.blacklistvalidator;

import java.util.List;

/**
 *
 * @author hcadavid
 */
public class Main {
    
    public static void main(String[] args){
        HostBlackListsValidator hblv=new HostBlackListsValidator();
        int n = 202;
        List<Integer> blackListOcurrences=hblv.checkHost("212.24.24.55", n);
        System.out.println("The host was found in the following blacklists:"+blackListOcurrences);
        
    }
    
}
