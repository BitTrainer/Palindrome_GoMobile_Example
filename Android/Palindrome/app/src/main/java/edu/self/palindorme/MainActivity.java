package edu.self.palindorme;

import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.util.Log;
import android.view.View;
import android.widget.Button;
import android.widget.EditText;
import android.widget.TextView;
import self.edu.palindrome.check.PalindromeCheck;

public class MainActivity extends AppCompatActivity
                          implements View.OnClickListener{
    private EditText mInputEditText;
    private TextView mResultTextView;
    private Button mCheckButton;
    private static String IS_PALINDROME_RESPONSE,
                          IS_NOT_PALINDROME_RESPONSE;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
        IS_PALINDROME_RESPONSE = getString(R.string.is_palindrom_response);
        IS_NOT_PALINDROME_RESPONSE = getString(R.string.is_not_a_palindrom_response);

        mInputEditText = (EditText) findViewById(R.id.txtInput);
        mResultTextView = (TextView) findViewById(R.id.txtResult);
        mCheckButton = (Button) findViewById(R.id.btnCheck);

        mCheckButton.setOnClickListener(this);
    }

    @Override
    public void onClick(View v) {
        String inputText = mInputEditText.getText().toString();
        boolean isPalindrome = PalindromeCheck.IsPalindrome(inputText);
        mResultTextView.setText(isPalindrome?IS_PALINDROME_RESPONSE:
                                             IS_NOT_PALINDROME_RESPONSE);

    }
}
