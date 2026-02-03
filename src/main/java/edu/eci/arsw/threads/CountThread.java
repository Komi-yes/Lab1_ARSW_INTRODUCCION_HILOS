/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package edu.eci.arsw.threads;

/**
 * Thread implementation that counts from a minimum to a maximum value.
 * @author Daniel Palacios Moreno
 */
public class CountThread extends  Thread{
    private int min;
    private int max;

    public CountThread (int min , int max ){
        this.min = min;
        this.max = max;
    }

    public void run(){
        for(int i = this.min; i <= this.max; i++){
            System.out.println(i);
        }
    }
}
