<?php
$file = fopen("serverList.txt", "r") or exit("Unable to open file!");
//Output a line of the file until the end is reached
while(!feof($file))
  {

  $lines = fgets($file);
  $eachVar = explode(" ", $lines);
  //PropertyTag
  $linestrtag = '<a href=http://'.$eachVar[0].':8003'.'>'.$eachVar[2].'</a><br>';
  echo $linestrtag;
  }
fclose($file);
?>
