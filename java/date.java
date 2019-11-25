import java.util.Date;
import java.text.SimpleDateFormat;
import java.util.Calendar;

class DateTest {
	public static void main(String[] args) {
                calendar();
        }
        public static void date(){
                Date now= new Date();
                //打印当前时间
                System.out.println("当前时间:"+now);
                System.out.println("当前时间toString():"+now.toString());
                System.out.println("当前时间getTime()："+now.getTime());
                /*
                当前时间:Mon Nov 25 15:15:14 CST 2019
                当前时间toString():Mon Nov 25 15:15:14 CST 2019
                当前时间getTime()：1574666114570
                */
                Date zero = new Date(0);
                System.out.println("用0作为构造方法，得到的日期是:"+zero);//用0作为构造方法，得到的日期是:Thu Jan 01 08:00:00 CST 1970
                        
                SimpleDateFormat sdf = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss SSS" );//H-代表24进制的小时 h-代表12进制的小时 S-代表毫秒
                Date d = new Date();
                String str = sdf.format(d);
                System.out.println("当前时间通过 yyyy-MM-dd HH:mm:ss SSS 格式化后的输出: "+str);//当前时间通过 yyyy-MM-dd HH:mm:ss SSS 格式化后的输出: 2019-11-25 15:07:57

                //创建一个从1995.1.1 00:00:00 到 1995.12.31 23:59:59 之间的随机日期
                sdf = new SimpleDateFormat("yyyy.MM.dd HH:mm:ss");
                String startStr = "1995.1.1 00:00:00";
                String endStr = "1995.12.31 23:59:59";
                
                try{
                        Date start = sdf.parse(startStr); //字符串转日期
                        Date end = sdf.parse(endStr);
                        long random = (long)(Math.random() * (end.getTime() - start.getTime()) + start.getTime());
                        System.out.println(new Date(random));//Tue Sep 26 04:58:33 CST 1995
                }catch(Exception e){
                        System.out.println(e);
                }
        }

        public static void calendar(){
                //采用单例模式获取日历对象Calendar.getInstance();
                Calendar c = Calendar.getInstance();

                //通过日历对象得到日期对象
                Date d = c.getTime();
          
                Date d2 = new Date(0);
                c.setTime(d2); //把这个日历，调成日期 : 1970.1.1 08:00:00
                System.out.println(c);

                c = Calendar.getInstance();
                Date now = c.getTime();
                // 当前日期
                System.out.println("当前日期：\t" + format(c.getTime()));

                // 下个月的今天
                c.setTime(now);
                c.add(Calendar.MONTH, 1);
                System.out.println("下个月的今天:\t" +format(c.getTime()));

                // 去年的今天
                c.setTime(now);
                c.add(Calendar.YEAR, -1);
                System.out.println("去年的今天:\t" +format(c.getTime()));

                // 上个月的第三天
                c.setTime(now);
                c.add(Calendar.MONTH, -1);
                c.set(Calendar.DATE, 3);
                System.out.println("上个月的第三天:\t" +format(c.getTime()));

                /*
                当前日期：      2019-11-25 16:08:26
                下个月的今天:   2019-12-25 16:08:26
                去年的今天:     2018-11-25 16:08:26
                上个月的第三天: 2019-10-03 16:08:26
                */

                // 下个月的倒数第三天
                c = Calendar.getInstance();//日历类采用了单例模式
                c.add(Calendar.MONTH, 1);
                int month = c.get(Calendar.MONTH);
                while (c.get(Calendar.MONTH) == month) {
                    c.add(Calendar.DATE, 1);
                    System.out.println("c.add(DATE,1)结果：" + sdf.format(c.getTime()));
                }
                c.add(Calendar.DATE, -3);
                System.out.println("下个月的倒数第三天：" + sdf.format(c.getTime()));
                /*
                c.add(DATE,1)结果：2019-12-26 16:12:43
                c.add(DATE,1)结果：2019-12-27 16:12:43
                c.add(DATE,1)结果：2019-12-28 16:12:43
                c.add(DATE,1)结果：2019-12-29 16:12:43
                c.add(DATE,1)结果：2019-12-30 16:12:43
                c.add(DATE,1)结果：2019-12-31 16:12:43
                c.add(DATE,1)结果：2020-01-01 16:12:43
                下个月的倒数第三天：2019-12-29 16:12:43
                */
        }

        private static SimpleDateFormat sdf = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss");
        private static String format(Date time) {
                return sdf.format(time);
        }
}

/*





*/