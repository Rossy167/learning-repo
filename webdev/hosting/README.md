# proper hosting not apache on a pi
I own rossy167.tech on namecheap, changed DNS from namecheap's to Amazons and pointed them at an S3 bucket

hosted on:
    ns-1570.awsdns-04.co.uk
    ns-781.awsdns-33.net
    ns-212.awsdns-26.com
    ns-1132.awsdns-13.org

Turns out that won't work if I want a HTTPS connection, looking into changing to CloudFront instead.

I used this guide to host the site on AWS, doesn't cost much and is a technology that's important in this industry:
https://benjamincongdon.me/blog/2017/06/13/How-to-Deploy-a-Secure-Static-Site-to-AWS-with-S3-and-CloudFront/
