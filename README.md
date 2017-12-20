#or http://176.9.3.149:14265
# 1.4.1.2
# 1.4.1.4

curl -H "Content-Type: application/json" -H "X-IOTA-API-Version: 1" -d '{"command": "getNodeInfo"}' http://iota.bitfinex.com
curl -H "Content-Type: application/json" -H "X-IOTA-API-Version: 1" -d '{"command": "getNodeInfo"}' http://176.9.3.149:14265

curl -H "Content-Type: application/json" -H "X-IOTA-API-Version: 1" -d '{"command": "getTrytes", "hashes": ["TGAGSHUKIDZLFXTXMDCAOAWDVFCOXFGXYCJANJRFWKDGCWCJHQNLMLKVOLRWUHECBVNEXA9VBHEJ99999","NEFCUE9SSXACIAOEHHEBRHOP9LMUQEBRBQDEJXREYDYN9YIHJSMHVNFLPQMMABGALXYBFIBD9ZUEZ9999","9AZYTTNJOMRZTYCHSP9AGSCVISZRPEFBLMRZNRNRWYDSIXMSWVSJDGFWXHTCVFJVP9IBEZQSWWZMZ9999"]}' http://iota.bitfinex.com
curl -H "Content-Type: application/json" -H "X-IOTA-API-Version: 1" -d '{"command": "getTrytes", "hashes": ["NEFCUE9SSXACIAOEHHEBRHOP9LMUQEBRBQDEJXREYDYN9YIHJSMHVNFLPQMMABGALXYBFIBD9ZUEZ9999"]}' http://176.9.3.149:14265

curl -H "Content-Type: application/json" -H "X-IOTA-API-Version: 1" -d '{"command": "getInclusionStates", "transactions": ["XHTZXCNGJHMYDSXPSQWOCEOCUZNWWQIZSBHAIHMFERFDUC9XFEUMDXVIWJLDXJSZQSLXJWPTBPUWTLNCUUUIDXUZYVIFJJ9K9BASGRXQQOXVXDSAVAZLWJFLRQKAFRRZ9UZRGVPSLRDJRFFYAJHWFHABECCOJFOFPVHNPXYVTJYPSH9SCBVKMSKKZGASWNSN9ANJO9GEMCLDDYNJECHCBCGEUDECOHCFNGEWADSUFHWG9YYLEVWSYYKBTPGQSEBPZBZU9XJLMNFIHNZCNONBICTJGCISSTOANCUNZLJDFTEVJTDHMUKLJQIGAWBAVKWPYKOHJCBFGTBK9MFIGDPVSBJNSLNWYKGMDICD9IVVXX9MAICOFIEGZPVQZMTTOTEDIXZDX9KODCIIUCTQSJ9IHGXNWSFSBYHZQ9BLUQRGNBDIDTUXBUBQKQHPIQRP9RHLJBVBNZNUCFAUUS9KG9KHXECJLRASWFLPPDHWDVIKHAGBGIO9RTZESYNTBDQQGKZWAXYIGVLVHPDLRPZFPXYKL9KL9LIKTIWUHGFYGSUO9EKCQKMGLJRBMZZ9YUUP99LDIHUWURORBEWUSZ9BLBQHYKRDNWCKNGCJTDFHBXHEIGJTZOJRKAMDDEAKTAERVYPSFHTCWFAVQJHUDWLPDWGCOGQRUNBHHMJWQBFSGGSDUQXHRQNXVWGGEFWACPSJQKILVOPUGDOFIZPXNWLADWQYK9TOOBFLFZMFLAJPACSEYMLCNCIVXJEVI9PZSSMDSZNPXGH9RMUCUXX9GKYV9CDJDDHOPFEDDMAFOIIKANGNVVMEOAO9ZCPRCODWXFVVRGUNRTKIZYXZEXAATCDJEQZBPMJFUDDUDRPNBJNRZFQQHNFBSKMPP9THDLABWSUMTHHH9LCAZDOIZJSGBJ9STEOQ9VEYKFZXMWSBRGTSYWDJBLZBMEDZCOQM9PUNBFIGH9BEPDMKIVPACZMRXDQTXWHWCHKQJGJTEFUQFOPAUELCUPYMAEAYNAOTFBGZYH9RDXCNWXKPRILYZXQFSCYUXKS9PAAXB9KBMEOMETIZKQORCYKOEWXWDOZKGNCBZIGEOCPHNC9GPIJOESJWMVKMTSISA9PK9SFBHDDNAYKKHQLNUVUWU9ECAICUYR9COEBWXAVI9YEKF9YEOFZSVXKTJHMDMQULXAXPXOCXCZHGCCNAKOBWNAVIPJPIXYWWZFFZPYJOECOHH9LGFVMWSIFFWLTEWTNEWSJZFVN99ZUBJNHUVOMRHAAWSOFBDJSFWNPCGNSASYDZVBGRYCRFTIEGI9TESGBRFYFC9DDSA9R9WNKCZZEDRBITSPKDMWMPBTPDMTNYAOPMH9OGRWBGYYNBMGODPIQAQVDKDFHOL9PPZXUCVVRWAUYTWUFDEEWYIZPN9OXYUIALTBGYAMKIJCQFDZKCZQTKUJWUYKBHVANAYALFMYPWHUUXBPWIQKJGBHVUFCXRMJGDXWPWAGHQEIILVUZNUEEFCCTKOFLJGLXIREOZPMVOCULKFVLXHZPYMRZBQUKJXOWHHTBQSIDQLYHGXYILWINYIFDJU9JWEFFOOEBZCZEPQYQRUYFSEPCSRZRJXJAXDKYOUPOQUSEWGUHWJEYDKMWKGETWCMNPVSEDH9MYXDAWYCULIQDATCIW9ERCEROCMLMELWNPOABULBUKGEZTQHBNUEUJZDLGEGVQUHFUPCVTKWTJ9JNNDRJSUCFLZRDOICF9OPXAOBQVKZJEKY9XTFIULUITZPVQMGHLTKMMVRABCMSGNGHWVPBZRCYBSDWIDDSGZFBKTSNCCBLQSEBEDRMAZEOCTPNZZDCYUKDTYFEIWYXIBUJJKCMQIKBSQXYPHVCJLJKVICUGNAYP9X9LFECELEFNWAQNNNKBAGDOHICFDZNMVJVZVGUBVDFRDSKMTPUFLGGQFQWQLPNDPSLGFYAALLDGPZLIFAVIEOQDBVFVFVBQVYPKBTIAEBAKZEOSPAPNIAARWIFHMKKTBQCBZUW9TENEZBBBZPXTVUBWO9X9ZDFITQSY9VO9WNJPKSDGIXIYFBUZOMSWHPGJYUSELWSIGLMQYODWPQNBYUJVBOGBYTBBLSAMALPEOJPCXXCBCEEQIHEEYIZTKWVNUCIGYNCZUULYFWOIJSVHCBOGSISFLSPSJXIJMSXZXRIIFFA9VJH9ORSGCW9KPWCHICGJZXKE9GSUDXZYUAPLHAKAHYHDXNPHENTERYMMBQOPSQIDENXKLKCEYCPVTZQLEEJVYJZV9BWU999999999999999999999999999ACPPA9999999999999999999999XZYIMXD99999999999A99999999ZWAKVBXHGXRLORLDTBMMBOGWZZPOGXXGR9PUMWBPHACLJKZLPSYFBQQQSIDFYINAIUSTRFBQMOHTJWNZWYSMZSWDKKOTDEXOKYYQLDKLRBCVNFSYCSQFCEQPREXYKDXYRPOQTUFPARKQPHPPVQHJXCTIJ9IEKA9999NMBUPOFMRYYLWKZPAFBYRCUVRVGPIMSJDGJJLAZOKFAFRIQINEBBCAQSXII9IDLSRUSEOOZQGIZGA9999999999999999999999999999999999999999999999999999999999F9RVUQCFRMXQEEEWLJSPMAICEIC"]}' http://iota.bitfinex.com

R9A9SDIVEENLHXNBRSHNTDMARBEOESYJR9HTDRYJAQLGOOQIPUVYTQOWY9OMAZYBHHNYBFLAWB9JA9999




{"trytes":["XHTZXCNGJHMYDSXPSQWOCEOCUZNWWQIZSBHAIHMFERFDUC9XFEUMDXVIWJLDXJSZQSLXJWPTBPUWTLNCUUUIDXUZYVIFJJ9K9BASGRXQQOXVXDSAVAZLWJFLRQKAFRRZ9UZRGVPSLRDJRFFYAJHWFHABECCOJFOFPVHNPXYVTJYPSH9SCBVKMSKKZGASWNSN9ANJO9GEMCLDDYNJECHCBCGEUDECOHCFNGEWADSUFHWG9YYLEVWSYYKBTPGQSEBPZBZU9XJLMNFIHNZCNONBICTJGCISSTOANCUNZLJDFTEVJTDHMUKLJQIGAWBAVKWPYKOHJCBFGTBK9MFIGDPVSBJNSLNWYKGMDICD9IVVXX9MAICOFIEGZPVQZMTTOTEDIXZDX9KODCIIUCTQSJ9IHGXNWSFSBYHZQ9BLUQRGNBDIDTUXBUBQKQHPIQRP9RHLJBVBNZNUCFAUUS9KG9KHXECJLRASWFLPPDHWDVIKHAGBGIO9RTZESYNTBDQQGKZWAXYIGVLVHPDLRPZFPXYKL9KL9LIKTIWUHGFYGSUO9EKCQKMGLJRBMZZ9YUUP99LDIHUWURORBEWUSZ9BLBQHYKRDNWCKNGCJTDFHBXHEIGJTZOJRKAMDDEAKTAERVYPSFHTCWFAVQJHUDWLPDWGCOGQRUNBHHMJWQBFSGGSDUQXHRQNXVWGGEFWACPSJQKILVOPUGDOFIZPXNWLADWQYK9TOOBFLFZMFLAJPACSEYMLCNCIVXJEVI9PZSSMDSZNPXGH9RMUCUXX9GKYV9CDJDDHOPFEDDMAFOIIKANGNVVMEOAO9ZCPRCODWXFVVRGUNRTKIZYXZEXAATCDJEQZBPMJFUDDUDRPNBJNRZFQQHNFBSKMPP9THDLABWSUMTHHH9LCAZDOIZJSGBJ9STEOQ9VEYKFZXMWSBRGTSYWDJBLZBMEDZCOQM9PUNBFIGH9BEPDMKIVPACZMRXDQTXWHWCHKQJGJTEFUQFOPAUELCUPYMAEAYNAOTFBGZYH9RDXCNWXKPRILYZXQFSCYUXKS9PAAXB9KBMEOMETIZKQORCYKOEWXWDOZKGNCBZIGEOCPHNC9GPIJOESJWMVKMTSISA9PK9SFBHDDNAYKKHQLNUVUWU9ECAICUYR9COEBWXAVI9YEKF9YEOFZSVXKTJHMDMQULXAXPXOCXCZHGCCNAKOBWNAVIPJPIXYWWZFFZPYJOECOHH9LGFVMWSIFFWLTEWTNEWSJZFVN99ZUBJNHUVOMRHAAWSOFBDJSFWNPCGNSASYDZVBGRYCRFTIEGI9TESGBRFYFC9DDSA9R9WNKCZZEDRBITSPKDMWMPBTPDMTNYAOPMH9OGRWBGYYNBMGODPIQAQVDKDFHOL9PPZXUCVVRWAUYTWUFDEEWYIZPN9OXYUIALTBGYAMKIJCQFDZKCZQTKUJWUYKBHVANAYALFMYPWHUUXBPWIQKJGBHVUFCXRMJGDXWPWAGHQEIILVUZNUEEFCCTKOFLJGLXIREOZPMVOCULKFVLXHZPYMRZBQUKJXOWHHTBQSIDQLYHGXYILWINYIFDJU9JWEFFOOEBZCZEPQYQRUYFSEPCSRZRJXJAXDKYOUPOQUSEWGUHWJEYDKMWKGETWCMNPVSEDH9MYXDAWYCULIQDATCIW9ERCEROCMLMELWNPOABULBUKGEZTQHBNUEUJZDLGEGVQUHFUPCVTKWTJ9JNNDRJSUCFLZRDOICF9OPXAOBQVKZJEKY9XTFIULUITZPVQMGHLTKMMVRABCMSGNGHWVPBZRCYBSDWIDDSGZFBKTSNCCBLQSEBEDRMAZEOCTPNZZDCYUKDTYFEIWYXIBUJJKCMQIKBSQXYPHVCJLJKVICUGNAYP9X9LFECELEFNWAQNNNKBAGDOHICFDZNMVJVZVGUBVDFRDSKMTPUFLGGQFQWQLPNDPSLGFYAALLDGPZLIFAVIEOQDBVFVFVBQVYPKBTIAEBAKZEOSPAPNIAARWIFHMKKTBQCBZUW9TENEZBBBZPXTVUBWO9X9ZDFITQSY9VO9WNJPKSDGIXIYFBUZOMSWHPGJYUSELWSIGLMQYODWPQNBYUJVBOGBYTBBLSAMALPEOJPCXXCBCEEQIHEEYIZTKWVNUCIGYNCZUULYFWOIJSVHCBOGSISFLSPSJXIJMSXZXRIIFFA9VJH9ORSGCW9KPWCHICGJZXKE9GSUDXZYUAPLHAKAHYHDXNPHENTERYMMBQOPSQIDENXKLKCEYCPVTZQLEEJVYJZV9BWU999999999999999999999999999ACPPA9999999999999999999999XZYIMXD99999999999A99999999ZWAKVBXHGXRLORLDTBMMBOGWZZPOGXXGR9PUMWBPHACLJKZLPSYFBQQQSIDFYINAIUSTRFBQMOHTJWNZWYSMZSWDKKOTDEXOKYYQLDKLRBCVNFSYCSQFCEQPREXYKDXYRPOQTUFPARKQPHPPVQHJXCTIJ9IEKA9999NMBUPOFMRYYLWKZPAFBYRCUVRVGPIMSJDGJJLAZOKFAFRIQINEBBCAQSXII9IDLSRUSEOOZQGIZGA9999999999999999999999999999999999999999999999999999999999F9RVUQCFRMXQEEEWLJSPMAICEIC"]















1. is a hash decodeable to show more info like value of transaction?
2. rockdb

1. Make a new address
2. Send money to existing address
3. check balance of an address
4. read a chain if possible back 10 hops
5. find genisis addresses
6. recreate gensis seed locally

curl -H "Content-Type: application/json" -H "X-IOTA-API-Version: 1" -d '{"command": "getBalances", "addresses": ["Z99ESC9LPRIPEPONNQTNMNGZREEEUWCREDYBSCMZHKZYWSACLZVQHZSARUIFGLGZKBPN9ITFIFJTZ9999"], "threshold": 100}' http://176.9.3.149:14265

cat /dev/urandom |LC_ALL=C tr -dc 'A-Z9' | fold -w 81 | head -n 1

# Z99ESC9LPRIPEPONNQTNMNGZREEEUWCREDYBSCMZHKZYWSACLZVQHZSARUIFGLGZKBPN9ITFIFJTZ9999
# JOCWIPIVIQOYMXOGZOWNODGLNWWSIP9XUFMFUKKLTSIJ9JVMPW9KHTL9PTAOKBFEU9MGTVLUHDIG99999
# GMJZ9FLELGXVPOHFCOIOOAIQCU9KOWOZSYBKXHRYLZQWDQUXTZLOS9XBHTQGURNGT9LIKPOLSUKLA9999
# AVEZQIWNRJJIWTYPTNCROHTENYFB9CXQBWDMU9UNIMIUQTYKEGM9GFNYLVYDUJ9EOSDFCZQLFDGW99999

udp://iota2.dedyn.io:14600 udp://tol06-h01-176-135-173-175.dsl.sta.abo.bbox.fr:14600 udp://io.iota.community:14600 udp://node7.iota.community:14600 udp://sync.iotasupport.com:14265 udp://gernat.ddns.net:14265 udp://static.224.135.130.94.clients.your-server.de:14265


 curl -H "Content-Type: application/json" -H "X-IOTA-API-Version: 1" -d '{"command": "findTransactions", "bundles": ["JNHOAP9YQQPCVPLJXBTTOHWAPSYVGGKFHROHPPHJXZYXXXUGFXSDOKTYDTIQCMMPT9HLLTNVVMRWVC9VY"]}' http://mainnet.necropaz.com:14500
 curl -H "Content-Type: application/json" -H "X-IOTA-API-Version: 1" -d '{"command": "findTransactions", "approvees": ["RDWRJFUSUOSYQZIIZSEBFFELUFVFVZKWMCFXUBVYIBGBILNZHKIJXB9HALEAACSVFZRNCZNZAGQJZ9999"]}' http://mainnet.necropaz.com:14500
{"hashes":["YCVXV9INXSINTGGOAWKNC9AXDAMUITHNCAQKCTTIVVSGXHLDHHVLGRDWZLGOCWSPZXNBKSS9ANRRA9999"],"duration":22}~/srcmand": "findTransactions", "approvees": ["YCVXV9INXSINTGGOAWKNC9AXDAMUITHNCAQKCTTIVVSGXHLDHHVLGRDWZLGOCWSPZXNBKSS9ANRRA9999"]}' http://mainnet.necropaz.com:14500
{"hashes":["AAIAIQKJOEGCZGTHI9KX9RMVPCPAUR9MYGFQKBAPBQJTQTUGLAMFV9IAB9XAOYRRVLWJDLVHGRHKZ9999"],"duration":2}~/src/github.com/andrewarrow/iotame $ 
~/src/github.com/andrewarrow/iotame $ 
~/src/github.com/andrewarrow/iotame $ curl -H "Content-Type: application/json" -H "X-IOTA-API-Version: 1" -d '{"command": "findTransactions", "approvees": ["AAIAIQKJOEGCZGTHI9KX9RMVPCPAUR9MYGFQKBAPBQJTQTUGLAMFV9IAB9XAOYRRVLWJDLVHGRHKZ9999"]}' http://mainnet.necropaz.com:14500
{"hashes":["GEPV9WYJNWLJJCTFE9QY9BODACKJIHVNJIVSXAZS9DAKDKASWZAMREHCEWTUCCROBR9SYGYANMVL99999"],"duration":1}


curl -H "Content-Type: application/json" -H "X-IOTA-API-Version: 1" -d '{"command": "getNeighbors"}' http://mainnet.necropaz.com:14500
{"neighbors":[{"address":"static.224.135.130.94.clients.your-server.de:14265","numberOfAllTransactions":3250,"numberOfRandomTransactionRequests":192,"numberOfNewTransactions":1195,"numberOfInvalidTransactions":0,"numberOfSentTransactions":64342,"connectionType":"udp"},{"address":"alon-e.mooo.com:15600","numberOfAllTransactions":84535,"numberOfRandomTransactionRequests":6043,"numberOfNewTransactions":105,"numberOfInvalidTransactions":0,"numberOfSentTransactions":110512,"connectionType":"tcp"},{"address":"gernat.ddns.net:14265","numberOfAllTransactions":0,"numberOfRandomTransactionRequests":0,"numberOfNewTransactions":0,"numberOfInvalidTransactions":0,"numberOfSentTransactions":64153,"connectionType":"udp"},{"address":"sync.iotasupport.com:14265","numberOfAllTransactions":9298,"numberOfRandomTransactionRequests":6000,"numberOfNewTransactions":0,"numberOfInvalidTransactions":0,"numberOfSentTransactions":70150,"connectionType":"udp"},{"address":"node7.iota.community:14600","numberOfAllTransactions":71393,"numberOfRandomTransactionRequests":6101,"numberOfNewTransactions":1268,"numberOfInvalidTransactions":0,"numberOfSentTransactions":70376,"connectionType":"udp"},{"address":"io.iota.community:14600","numberOfAllTransactions":0,"numberOfRandomTransactionRequests":0,"numberOfNewTransactions":0,"numberOfInvalidTransactions":0,"numberOfSentTransactions":64153,"connectionType":"udp"},{"address":"tol06-h01-176-135-173-175.dsl.sta.abo.bbox.fr:14600","numberOfAllTransactions":70230,"numberOfRandomTransactionRequests":6221,"numberOfNewTransactions":9660,"numberOfInvalidTransactions":0,"numberOfSentTransactions":100172,"connectionType":"udp"},{"address":"iota2.dedyn.io:14600","numberOfAllTransactions":109664,"numberOfRandomTransactionRequests":9369,"numberOfNewTransactions":41098,"numberOfInvalidTransactions":0,"numberOfSentTransactions":86955,"connectionType":"udp"}],"duration":1}
