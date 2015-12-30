//
//  ViewController.swift
//  IsItAPalindrome
//
//  Created by Grant van Staden on 29/12/2015.
//  Copyright © 2015 Grant van Staden. All rights reserved.
//

import UIKit
import PalindromeCheck

class ViewController: UIViewController {

    @IBOutlet weak var _inputTextField: UITextField!
    @IBOutlet weak var _outputLabel: UILabel!
    
    
    override func viewDidLoad() {
        super.viewDidLoad()
        // Do any additional setup after loading the view, typically from a nib.
    }

    override func didReceiveMemoryWarning() {
        super.didReceiveMemoryWarning()
        // Dispose of any resources that can be recreated.
    }

    @IBAction func OnCheckClick(sender: UIButton) {
        
       let candidate = _inputTextField.text!
       let isPalindrome = GoPalindromeCheckIsPalindrome(candidate);
        if (isPalindrome) {
                _outputLabel.text = "Yes, this is a palindrome."
        } else {
            _outputLabel.text = "No, this is not a palindrome."
        }
        
    }

}

