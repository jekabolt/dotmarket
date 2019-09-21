window.addEventListener("load", function() {
    // store tabs variable
    var myTabs = document.querySelectorAll("ul.nav-tabs > li");

    function myTabClicks(tabClickEvent) {
        for (var i = 0; i < myTabs.length; i++) {
            myTabs[i].classList.remove("active");
        }
        var clickedTab = tabClickEvent.currentTarget;
        clickedTab.classList.add("active");
        tabClickEvent.preventDefault();
        var myContentPanes = document.querySelectorAll(".tab-pane");
        for (i = 0; i < myContentPanes.length; i++) {
            myContentPanes[i].classList.remove("active");
        }
        var anchorReference = tabClickEvent.target;
        var activePaneId = anchorReference.getAttribute("href");
        var activePane = document.querySelector(activePaneId);
        activePane.classList.add("active");
    }
    for (i = 0; i < myTabs.length; i++) {
        myTabs[i].addEventListener("click", myTabClicks)
    }
});
window.addEventListener("load", loadSuggests);


document.getElementById("suggests").addEventListener("click", loadSuggests);
document.getElementById("scheduled").addEventListener("click", loadScheduled);

function loadSuggests() {
    document.getElementById('showScrollsuggests').innerText = ""
    var xhr = new XMLHttpRequest();
    // if (document.getElementById('showScrollsuggests').innerText.length < 20) {

    xhr.onload = function() {
        if (xhr.status >= 200 && xhr.status < 300) {
            let resp = JSON.parse(xhr.responseText)
            resp.forEach(layout);

            function layout(item, index) {
                let images = ""

                function imagesLayout(item, index) {
                    images += ` <img src="` + item + `"
                    alt="">`
                }
                item.Images.forEach(imagesLayout);

                document.getElementById('showScrollsuggests').innerHTML += `
                                            <div class="receipt">
                <h1 class="logo">DOTMARKET/SUGGESTS</h1>
                <div class="address">
                    ISSUED : ` + timeStamp() + `
                </div>
                <div class="transactionDetails">
                    <div class="detail">Reg#17</div>
                    <div class="detail">TRN#1313</div>
                    <div class="detail">CSHR#00097655</div>
                    <div class="detail">str#9852</div>
                </div>
                <div class="transactionDetails">
                    DatePublished : ` + item.DatePublished + `
                </div>
                <div class="transactionDetails">
                    DateScheduled : ` + item.DateScheduled + `
                </div>
                <div class="transactionDetails">
                    Text : ` + item.Text + `
                </div>

                <div class="centerItem bold">
                    <div class="item">PostOwner : <a href="` + item.PostOwner + `">` + item.PostOwner + ` </a> </div>
                </div>


                <div class="receiptBarcode">
                    <div class="barcode">
                    ` + item.PostOwner + `
                    </div>
                </div>
                <div class="returnPolicy bold">
                RETURNS WITH RECEIPT THRU :  <br>
                    ` + timeStamp(3) + `
                </div>


                <div id="coupons" class="coupons">
                    <!--       start coupon -->
                    <div class="couponContainer">
                        <h1 class="logo">ATTACHMENTS</h1>

                       ` + images + `

                        <div class="expiration">
                            <div class="item bold">
                                Expires ` + timeStamp(3) + `
                            </div>
                        </div>
                        <div class="barcode">
                           vk.com/dotmarket
                        </div>
                    </div>
                    <!--       end coupon -->
                </div>
            </div>`;
            }
        } else {
            console.log('The request failed!');
        }
    };

    xhr.open('GET', 'http://157.245.84.45/api/v1.0/suggests');
    xhr.send();
    // }
}

function loadScheduled() {
    document.getElementById('showScrollscheduled').innerText = ""

    // if (document.getElementById('showScrollscheduled').innerText.length < 20) {
    var xhr = new XMLHttpRequest();

    xhr.onload = function() {
        if (xhr.status >= 200 && xhr.status < 300) {
            let resp = JSON.parse(xhr.responseText)
            resp.forEach(layout);

            function layout(item, index) {
                let images = ""

                function imagesLayout(item, index) {
                    images += ` <img src="` + item + `"
                    alt="">`
                }
                item.Images.forEach(imagesLayout);

                document.getElementById('showScrollscheduled').innerHTML += `
                                            <div class="receipt">
                <h1 class="logo">DOTMARKET/SUGGESTS</h1>
                <div class="address">
                    ISSUED : ` + timeStamp() + `
                </div>
                <div class="transactionDetails">
                    <div class="detail">Reg#17</div>
                    <div class="detail">TRN#1313</div>
                    <div class="detail">CSHR#00097655</div>
                    <div class="detail">str#9852</div>
                </div>
                <div class="transactionDetails">
                    DatePublished : ` + item.DatePublished + `
                </div>
                <div class="transactionDetails">
                    DateScheduled : ` + item.DateScheduled + `
                </div>
                <div class="transactionDetails">
                    Text : ` + item.Text + `
                </div>

                <div class="centerItem bold">
                    <div class="item">PostOwner : <a href="` + item.PostOwner + `">` + item.PostOwner + ` </a> </div>
                </div>


                <div class="receiptBarcode">
                    <div class="barcode">
                    ` + item.PostOwner + `
                    </div>
                </div>
                <div class="returnPolicy bold">
                RETURNS WITH RECEIPT THRU :  <br>
                    ` + timeStamp(3) + `
                </div>


                <div id="coupons" class="coupons">
                    <!--       start coupon -->
                    <div class="couponContainer">
                        <h1 class="logo">ATTACHMENTS</h1>

                       ` + images + `

                        <div class="expiration">
                            <div class="item bold">
                                Expires ` + timeStamp(3) + `
                            </div>
                        </div>
                        <div class="barcode">
                           vk.com/dotmarket
                        </div>
                    </div>
                    <!--       end coupon -->
                </div>
            </div>`;
            }
        } else {
            console.log('The request failed!');
        }
    };

    xhr.open('GET', 'http://157.245.84.45/api/v1.0/postponed');
    xhr.send();
    // }
}

function timeStamp(days = 0) {
    // Create a date object with the current time
    var now = new Date();

    // Create an array with the current month, day and time
    var date = [now.getMonth() + 1, now.getDate() + days, now.getFullYear()];

    // Create an array with the current hour, minute and second
    var time = [now.getHours(), now.getMinutes(), now.getSeconds()];

    // Determine AM or PM suffix based on the hour
    var suffix = (time[0] < 12) ? "AM" : "PM";

    // Convert hour from military time
    time[0] = (time[0] < 12) ? time[0] : time[0] - 12;

    // If hour is 0, set it to 12
    time[0] = time[0] || 12;

    // If seconds and minutes are less than 10, add a zero
    for (var i = 1; i < 3; i++) {
        if (time[i] < 10) {
            time[i] = "0" + time[i];
        }
    }

    // Return the formatted string
    return date.join("/") + " " + time.join(":") + " " + suffix;
}