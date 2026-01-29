package com.example.app;

import android.os.Bundle;
import android.net.http.SslError;
import android.webkit.SslErrorHandler;
import android.webkit.WebView;
import com.getcapacitor.BridgeActivity;
import com.getcapacitor.BridgeWebViewClient;

public class MainActivity extends BridgeActivity {
    @Override
    public void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);

        // Enable WebView debugging for Chrome DevTools
        android.webkit.WebView.setWebContentsDebuggingEnabled(true);

        // WORKAROUND: Ignore SSL errors for local development
        WebView webView = getBridge().getWebView();
        webView.setWebViewClient(new BridgeWebViewClient(getBridge()) {
            @Override
            public void onReceivedSslError(WebView view, SslErrorHandler handler, SslError error) {
                // Ignore SSL certificate errors (for self-signed certs)
                handler.proceed();
            }
        });
    }
}
